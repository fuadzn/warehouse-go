pipeline {
  agent {
    docker {
      image 'golang:1.22-alpine'
    }
  }

  environment {
    CGO_ENABLED = '0'
    GOOS = 'linux'
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Dependencies') {
      steps {
        sh 'go mod tidy'
        sh 'go mod download'
      }
    }

    stage('Build') {
      steps {
        sh 'go build -o app ./...'
      }
    }

    stage('Test') {
      steps {
        sh 'go test ./...'
      }
    }
  }

  post {
    success {
      echo 'Build berhasil!'
    }
    failure {
      echo 'Build gagal!'
    }
  }
}