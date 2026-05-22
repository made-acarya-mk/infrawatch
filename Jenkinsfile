pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                echo 'Checking out source code...'
            }
        }

        stage('Run Go Tests') {
            steps {
                sh 'cd app && go test ./...'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t infrawatch-api ./app'
            }
        }

        stage('Run Container with Docker Compose') {
            steps {
                sh 'docker compose down || true'
                sh 'docker compose up -d --build'
            }
        }

        stage('Health Check') {
            steps {
                sh './scripts/health-check.sh'
            }
        }

        stage('Backup') {
            steps {
                sh './scripts/backup.sh'
            }
        }
    }

    post {
        success {
            echo 'Deployment completed successfully.'
        }

        failure {
            echo 'Deployment failed. Please check logs.'
        }
    }
}