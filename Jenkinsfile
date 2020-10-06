pipeline {
	agent none 
	environment { 
		SLACK_CHANNEL = '#studying' 
		REGISTRY = 'jisunpark/golang-sample' 
		REGISTRYCREDENTIAL = 'dockerhub' 
		DOCKER_IMAGE = ''
	}

	stages { 
		stage('Start') { 
			agent any 
			steps { 
				slackSend (channel: SLACK_CHANNEL, color: '#FFFF00', message: "STARTED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})") 
			} 
		}
		stage('Test') {
			agent { docker { image 'golang' } }

            steps {                                           
                // Create our project directory.
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/MY_PROJECT_DIRECTORY'

                // Copy all files in our Jenkins workspace to our project directory.                
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/MY_PROJECT_DIRECTORY'

                // Copy all files in our "vendor" folder to our "src" folder.
                sh 'cp -r ${WORKSPACE}/vendor/* ${GOPATH}/src'

                // Build the app.
                sh 'go build'               
            }
		}
		stage('Docker image build') { 
			agent any 
			when {
                branch 'master'
            }
            steps{
        		script {
          			DOCKER_IMAGE = docker.build registry + ":$BUILD_NUMBER"
        		}
      		}
		} 
		stage('Docker image push to Dockerhub') {
      		steps{
        		script {
          			docker.withRegistry('', REGISTRYCREDENTIAL) {
            			DOCKER_IMAGE.push()
          			}
        		}
      		}
    	}
		stage('Clean local docker image') { 
			agent any
				steps { 
				sh 'docker rmi $REGISTRY:$BUILD_NUMBER' 
			} 
		} 
	} 
	post { 
		success { 
			slackSend (channel: SLACK_CHANNEL, color: '#00FF00', message: "SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})") 
		} 
		failure { 
			slackSend (channel: SLACK_CHANNEL, color: '#FF0000', message: "FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})") 
		}
	} 
}
