<?php

function fibonacci($n)
{
  $f1 = 0;
  $f2 = 1;

  if ($n <= 0) {
    return null;
  }

  if ($n <= 2) {
    return $n - 1;
  }

  for ($i = 3; $i <= $n; $i++) {
    $fn = $f1 + $f2;
    $f1 = $f2;
    $f2 = $fn;
  }

  return $fn;
}
