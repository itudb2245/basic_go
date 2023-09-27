pipeline {
    agent any

    tools {
        go '1.21.1'
    }
    environment {
        GO111MODULE = 'on'
    }

    stages {
        stage('Build') {
            steps {
                echo 'Build Stage Started'
                checkout scm
                sh 'mkdir -p .build'
                //sh 'go mod init'
                sh 'go build main.go'
                //sh 'go mod init example.com/greetings'
                sh 'cp go.mod .build/go.mod.orig'
                echo 'Build Stage Ended'
            }
        }

        stage('Test') {
            steps {
                echo 'Test Stage Started'
                sh 'go test  main_test.go'
                echo 'Test Stage Ended'
            }
        }
    }
}