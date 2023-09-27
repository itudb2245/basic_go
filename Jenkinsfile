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
                checkout scm
                sh 'mkdir -p .build'
                sh 'cp go.mod .build/go.mod.orig'
            }
        }
    }
}