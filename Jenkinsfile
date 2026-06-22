pipeline {
  agent {
    docker {
        image 'golang:1.25-alpine'
    }
  }

  environment {
    CGO_ENABLED = '0'
    GOOS = 'linux'
    SERVICE = 'merchant-service'
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Dependencies') {
      steps {
            dir('merchant-service') {
                sh 'go mod tidy'
                sh 'go mod download'
            }
      }
    }

    stage('Build') {
      steps {
            dir('merchant-service') {
                sh 'go build -o app ./...'
            }
      }
    }

    stage('Test') {
      steps {
            dir('merchant-service') {
                sh 'go test ./...'
            }
      }
    }
  }

  post {
    success {
      echo 'Merchant service Build berhasil!'
    }
    failure {
      echo 'Merchant service Build gagal!'
    }
  }
}