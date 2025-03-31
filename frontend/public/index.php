<?php
require_once __DIR__ . '/../api/BankApiClient.php';

session_start();
$api = new BankApiClient('http://go-backend:8080');

// Handle login
if ($_SERVER['REQUEST_METHOD'] === 'POST' && isset($_POST['login'])) {
    $response = $api->login($_POST['username'], $_POST['password']);
    if ($response['success']) {
        $_SESSION['user'] = $response['user'];
        $_SESSION['token'] = $response['token'];
        header('Location: dashboard.php');
        exit;
    }
}

// Include template
include __DIR__ . '/../templates/header.php';
include __DIR__ . '/../templates/login.php';
include __DIR__ . '/../templates/footer.php';
?>
