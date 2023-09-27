pipeline {
    agent any

    tools {
        go '1.21.1'
    }
    environment {
        GO111MODULE = 'on'
    }

    stages {
        stage('Compile') {
            steps {
                sh 'go mod init'
                sh 'go build'
            }
        }
    }
}