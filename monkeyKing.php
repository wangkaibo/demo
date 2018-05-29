<?php

function monkeyKing($monkeys, $num)
{
  if (count($monkeys) === 0) {
    return null;
  }
  $i = 1;
  while(count($monkeys) > 1) {
    if ($i % $num === 0) {
      unset($monkeys[$i - 1]);
    } else {
      array_push($monkeys, $monkeys[$i - 1]);
      unset($monkeys[$i - 1]);
      var_dump($monkeys);
    }
    $i++;
  }
  return $monkeys; 
}

$arr = ['a', 'b', 'c', 'd', 'e', 'f', 'g'];

var_dump(monkeyKing($arr, 3));
