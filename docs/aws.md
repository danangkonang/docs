# PresignResponse.kt
```java
package com.humarizone.data.model.response

import com.google.gson.annotations.SerializedName
import kotlinx.serialization.Serializable

@Serializable
data class PresignResponse(
    val id: String,
    val url: String,
    @SerializedName("file_url")
    val fileUrl: String,
    val fields: Fields
)

@Serializable
data class Fields(
    @SerializedName("X-Amz-Algorithm")
    val xAmzAlgorithm: String,

    @SerializedName("X-Amz-Credential")
    val xAmzCredential: String,

    @SerializedName("X-Amz-Date")
    val xAmzDate: String,

    @SerializedName("X-Amz-Signature")
    val xAmzSignature: String,

    val key: String,

    val policy: String
)
```
# UploadService.kt
```java
package com.humarizone.data.service

import com.humarizone.data.model.ApiResponse
import com.humarizone.data.model.response.PresignResponse
import retrofit2.Response
import retrofit2.http.GET

interface UploadService {
    @GET("upload/presign-url")
    suspend fun presignUrl(): Response<ApiResponse<PresignResponse>>
}
```
# UploadRepository.kt
```java
package com.humarizone.data.repository

import com.google.gson.Gson
import com.humarizone.data.model.response.PresignResponse
import com.humarizone.data.service.UploadService
import com.humarizone.utils.Resource
import com.humarizone.utils.toResource
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class UploadRepository @Inject constructor(private val apiService: UploadService, private val gson: Gson) {
    suspend fun presignUrl(): Resource<PresignResponse> {
        return try {
            apiService.presignUrl().toResource(gson)
        } catch (e: Exception) {
            Resource.Error(e.message ?: "Koneksi bermasalah")
        }
    }
}
```
# AwsService.kt
```java
package com.humarizone.data.service

import com.humarizone.data.model.ApiResponse
import com.humarizone.data.model.response.PresignResponse
import okhttp3.MultipartBody
import okhttp3.RequestBody
import retrofit2.Response
import retrofit2.http.Multipart
import retrofit2.http.POST
import retrofit2.http.Part
import retrofit2.http.PartMap
import retrofit2.http.Url

interface AwsService {
    @Multipart
    @POST
    suspend fun upload(
        @Url url: String,
        @PartMap fields: Map<String, @JvmSuppressWildcards RequestBody>,
        @Part filePart: MultipartBody.Part
    ): Response<ApiResponse<PresignResponse>>
}
```
# AwsRepository.kt
```java
package com.humarizone.data.repository

import android.util.Log
import android.webkit.MimeTypeMap
import com.google.gson.Gson
import com.humarizone.data.model.ApiResponse
import com.humarizone.data.model.response.PresignResponse
import com.humarizone.data.service.AwsService
import com.humarizone.utils.Resource
import okhttp3.MultipartBody
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.asRequestBody
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.MediaType.Companion.toMediaTypeOrNull
import java.io.File
import javax.inject.Inject

class AwsRepository @Inject constructor(private val apiService: AwsService) {
    suspend fun upload(response: PresignResponse, file: File): Resource<PresignResponse> {
        return try {
            val extension = MimeTypeMap.getFileExtensionFromUrl(file.absolutePath) ?: MimeTypeMap.getFileExtensionFromUrl(file.name) ?: "jpg"
            val mimeType = MimeTypeMap.getSingleton().getMimeTypeFromExtension(extension) ?: "application/octet-stream"

            val fieldsMap:Map<String, RequestBody> = mapOf(
                "X-Amz-Algorithm" to response.fields.xAmzAlgorithm.toRequestBody("text/plain".toMediaTypeOrNull()),
                "X-Amz-Credential" to response.fields.xAmzCredential.toRequestBody("text/plain".toMediaTypeOrNull()),
                "X-Amz-Date" to response.fields.xAmzDate.toRequestBody("text/plain".toMediaTypeOrNull()),
                "X-Amz-Signature" to response.fields.xAmzSignature.toRequestBody("text/plain".toMediaTypeOrNull()),
                "key" to response.fields.key.toRequestBody("text/plain".toMediaTypeOrNull()),
                "acl" to "public-read".toRequestBody("text/plain".toMediaTypeOrNull()),
                "policy" to response.fields.policy.toRequestBody("text/plain".toMediaTypeOrNull()),
                "Content-Type" to mimeType.toRequestBody("text/plain".toMediaTypeOrNull())
            )

            val requestFile = file.asRequestBody(mimeType.toMediaTypeOrNull())
            val filePart = MultipartBody.Part.createFormData("file", file.name, requestFile)
            val response = apiService.upload(response.url,fieldsMap, filePart)
            Log.d("DANANG", response.code().toString())
            Log.d("DANANG", response.toString())
            if (response.isSuccessful) {
                Resource.Success(response.message())
            } else {
                val errorJson = response.errorBody()?.string()
                val errorResponse = try {
                    Gson().fromJson(errorJson, ApiResponse::class.java)
                } catch (_: Exception) {
                    null
                }
                Resource.Error(errorResponse?.message ?: response.message())
            }
        } catch (e: Exception) {
            Resource.Error(e.message ?: "Network error")
        }
    }
}
```