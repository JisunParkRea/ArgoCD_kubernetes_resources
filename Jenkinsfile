pipeline {
	agent none 

	environment { 
		SLACK_CHANNEL = '#jenkins_notification' 
		REGISTRY = '1223yys/spring-web-jenkins' 
		REGISTRYCREDENTIAL = '1223yys' 
	}

	stages { 
		stage('Start') { 
			agent any steps { 
				slackSend (channel: SLACK_CHANNEL, color: '#FFFF00', message: "STARTED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})") 
			} 
		} 
		stage('Test') { 
			agent { 
				docker { 
					image 'adoptopenjdk/openjdk13:alpine' 
					alwaysPull true 
				} 
			} 
			steps { 
				sh './gradlew test --no-daemon' 
			} 
		} 
		stage('Build') { 
			agent { 
				docker { 
					image 'adoptopenjdk/openjdk13:alpine' 
					alwaysPull true 
				} 
			} 
			steps { 
				sh './gradlew --no-daemon build -x test' 
			} 
			post { 
				success { 
					stash includes: '**/build/**', name: 'build' 
				} 
			} 
		} 
		stage('Docker image build & push') { 
			agent any steps { 
				unstash 'build' //docker username/password credential 
				sh 'docker login -u username -p password' 
				sh 'docker build -t $REGISTRY:latest .' 
			} 
			post { 
				success { 
					sh 'docker image ls | grep 1223yys' 
					sh 'docker push $REGISTRY:latest' 
				} 
			} 
		} 
		stage('Clean docker image') { 
			agent any steps { 
				sh 'docker rmi $REGISTRY' 
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
