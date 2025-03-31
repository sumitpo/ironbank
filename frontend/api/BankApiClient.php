<?php
class BankApiClient {
  private $baseUrl;

  public function __construct($baseUrl) {
    $this->baseUrl = $baseUrl;
  }

  private function callApi($endpoint, $method = 'GET', $data = []) {
    $url = $this->baseUrl . $endpoint;
    $options = [
      'http' => [
        'method'  => $method,
        'header'  => "Content-type: application/json\r\n",
        'content' => json_encode($data)
      ]
    ];

    $context = stream_context_create($options);
    $result = file_get_contents($url, false, $context);
    return json_decode($result, true);
  }

  public function login($username, $password) {
    return $this->callApi('/auth/login', 'POST', [
      'username' => $username,
      'password' => $password
    ]);
  }

  public function getAccounts($token) {
    return $this->callApi('/accounts', 'GET', [], [
      'Authorization: Bearer ' . $token
    ]);
  }

  // Add more API methods as needed
}
?>
