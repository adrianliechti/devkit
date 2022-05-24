import hudson.model.User
import hudson.security.HudsonPrivateSecurityRealm
import hudson.security.FullControlOnceLoggedInAuthorizationStrategy
import jenkins.model.Jenkins

def username = System.getenv("ADMIN_USERNAME")
def password = System.getenv("ADMIN_PASSWORD")

if (!username || username.allWhitespace) {
    return
}

if (!password || password.allWhitespace) {
    return
}

def realm = new HudsonPrivateSecurityRealm(false)
def strategy = new FullControlOnceLoggedInAuthorizationStrategy()

def user = User.get("admin", false)

if (user == null) {
    realm.createAccount(username, password)
}

strategy.setAllowAnonymousRead(false)

Jenkins.instance.setSecurityRealm(realm)
Jenkins.instance.setAuthorizationStrategy(strategy)
Jenkins.instance.save()