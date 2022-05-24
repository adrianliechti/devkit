import jenkins.model.Jenkins
import jenkins.model.JenkinsLocationConfiguration

def baseUrl = System.getenv("BASE_URL")

if (!baseUrl || baseUrl.allWhitespace) {
    return
}

def location = JenkinsLocationConfiguration.get()

location.setUrl(baseUrl)
location.save()