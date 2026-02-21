pipeline {
    agent any

    environment {
        // This matches your local build command name
        IMAGE_NAME = "apps.gin"
    }

    stages {
        stage('Checkout') {
            steps {
                // Pulls the latest code from GitHub
                checkout scm
            }
        }

        stage('Build Go App') {
            steps {
                script {
                    echo "Building the Go service..."
                    // This executes your specific build command
                    sh "docker build -t ${IMAGE_NAME} ."
                }
            }
        }

        stage('Verify Image') {
            steps {
                // Confirm the image was created successfully
                sh "docker images | grep ${IMAGE_NAME}"
            }
        }

        stage('Cleanup') {
            steps {
                // Removes 'dangling' images (old builds) to save space
                sh "docker image prune -f"
            }
        }
    }

    post {
        success {
            echo "Build successful! Image ${IMAGE_NAME} is ready."
        }
        failure {
            echo "Build failed. Check the logs above."
        }
    }
}
