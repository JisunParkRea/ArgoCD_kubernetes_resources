pipeline {
	agent none 

	environment { 
		SLACK_CHANNEL = '#studying' 
		REGISTRY = 'jisunpark/django-sample' 
		REGISTRYCREDENTIAL = '1223yys' 
	}

	stages { 
		stage('Start') { 
			agent any 
			steps { 
				slackSend (channel: SLACK_CHANNEL, color: '#FFFF00', message: "STARTED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})") 
			} 
		} 
		
	} 
}
