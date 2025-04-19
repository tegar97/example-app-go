pipeline {
    agent {
        dockerfile {
            filename 'Dockerfile'
            args '-v /var/run/docker.sock:/var/run/docker.sock'
        }
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o main .'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build -t hello-world-app:${BUILD_NUMBER} .'
            }
        }

        // Uncomment and configure this stage when you're ready to push to a Docker registry
        /*
        stage('Docker Push') {
            steps {
                withCredentials([string(credentialsId: 'docker-hub-credentials', variable: 'DOCKER_HUB_CREDENTIALS')]) {
                    sh 'docker login -u username -p ${DOCKER_HUB_CREDENTIALS}'
                    sh 'docker tag hello-world-app:${BUILD_NUMBER} username/hello-world-app:${BUILD_NUMBER}'
                    sh 'docker push username/hello-world-app:${BUILD_NUMBER}'
                }
            }
        }
        */
    }

    post {
        always {
            // Clean up
            sh 'docker rmi hello-world-app:${BUILD_NUMBER} || true'
        }
        success {
            echo 'Build completed successfully!'
        }
        failure {
            echo 'Build failed!'
        }
    }
}
