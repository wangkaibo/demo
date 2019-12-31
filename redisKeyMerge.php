<?php

$fp = fopen("./61500_no_num_sort.txt", "r");

$map = [];
$count = 0;
while(false !== ($line = fgets($fp))) {
    $line = preg_replace("/\s+/", ",", $line);
    $arr = explode(',', $line);

    $str = $arr[0];
    $m = $arr[1];
    $arr = explode(':', $str);
    

    $key = str_replace($arr[1], "", $str);

    if (isset($map[$key])) {
        $map[$key] += $m;
    } else {
        $map[$key] = $m;
    }
    $count++;
    echo $count . "\n";
}
$file = './redis_key_merge_' . date('Y-m-dHis') . '.csv';
if (file_exists($file)) {
    unlink($file);
}
$fileHandle = fopen($file, "a+");
arsort($map);
foreach ($map as $key => $size) {
    $line = [$key, $size];
    fwrite($fileHandle, implode(',', $line) . PHP_EOL);
}
fclose($fileHandle);