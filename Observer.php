<?php

interface IObserver
{
    function onChanged($sender, $args);
}

interface IObservable
{
    function addObserver($observer);
}

class UserList implements IObservable
{
    public $_observers = [];
    public $users = [];
    public function addUser($name) {
        $this->users[] = $name;
        foreach ($this->_observers as $ob) {
            $ob->onChanged($this, $name);
        }
    }

    public function addObserver($observer) {
        $this->_observers[] = $observer;
    }
}

class UserListLogger implements IObserver
{
    public function onChanged($sender, $args) {
        echo "observer print: user '{$args}' added";
    }
}

$userList = new UserList();
$userList->addObserver(new UserListLogger());
$userList->addUser('Pavel');
