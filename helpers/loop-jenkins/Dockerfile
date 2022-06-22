FROM openjdk:11-bullseye

RUN apt-get update -qq && apt-get install -qq -y --no-install-recommends \
    rsync vim \
    && rm -rf /var/lib/apt/lists/*

ENV LANG C.UTF-8

ENV XDG_RUNTIME_DIR /run/user/1000

ENV JENKINS_REF /usr/share/jenkins/ref
ENV JENKINS_HOME /var/jenkins_home

ARG JENKINS_VERSION
ENV JENKINS_VERSION ${JENKINS_VERSION:-2.346.1}

ARG JENKINS_PM_VERSION
ENV JENKINS_PM_VERSION ${JENKINS_PM_VERSION:-2.12.6}

ENV BASE_URL=http://localhost:8080

ENV ADMIN_USERNAME=admin
ENV ADMIN_PASSWORD=admin

ENV CASC_JENKINS_CONFIG=$JENKINS_HOME/jenkins.yaml

ENV JAVA_OPTS="-Djenkins.install.runSetupWizard=false -Dpermissive-script-security.enabled=true"

RUN mkdir -p $JENKINS_HOME $JENKINS_REF $XDG_RUNTIME_DIR \
    && chown 1000:1000 $JENKINS_HOME $JENKINS_REF $XDG_RUNTIME_DIR \
    && groupadd -g 1000 jenkins \
    && useradd -d "$JENKINS_HOME" -u 1000 -g 1000 -m -s /bin/bash jenkins

RUN curl -fsSL -o /usr/share/jenkins/jenkins.war https://repo.jenkins-ci.org/public/org/jenkins-ci/main/jenkins-war/${JENKINS_VERSION}/jenkins-war-${JENKINS_VERSION}.war
RUN curl -fsSL -o /usr/lib/jenkins-plugin-manager.jar https://github.com/jenkinsci/plugin-installation-manager-tool/releases/download/${JENKINS_PM_VERSION}/jenkins-plugin-manager-${JENKINS_PM_VERSION}.jar

RUN jar -xf /usr/share/jenkins/jenkins.war WEB-INF/lib/cli-${JENKINS_VERSION}.jar \
    && mv WEB-INF/lib/cli-${JENKINS_VERSION}.jar /usr/lib/jenkins-cli.jar \
    && rm -r WEB-INF

COPY jenkins /usr/local/bin/jenkins
COPY jenkins-cli /usr/local/bin/jenkins-cli
COPY jenkins-plugin-cli /usr/local/bin/jenkins-plugin-cli

COPY --chown=1000:1000 ref/ $JENKINS_REF/

RUN jenkins-plugin-cli --plugin-file $JENKINS_REF/plugins.txt

# Docker CLI
ENV DOCKER_CLIENT_VERSION="20.10.16"
RUN curl -fsSL "https://download.docker.com/linux/static/stable/$(uname -m)/docker-${DOCKER_CLIENT_VERSION}.tgz" | tar -zxf - --strip=1 -C /usr/local/bin/ docker/docker

USER 1000

EXPOSE 8080
EXPOSE 50000

VOLUME $JENKINS_HOME

CMD [ "/usr/local/bin/jenkins" ]