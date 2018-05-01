#!/usr/bin/env bash
set -ex

# https://maven.apache.org/guides/getting-started/maven-in-five-minutes.html
generateMavenProject(){
	mvn archetype:generate -DgroupId=com.mycompany.app -DartifactId=my-app -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
}

build(){
	rm -rf target
	mvn package
}

run(){
	cat 4_big.txt | java -cp target/my-app-1.0-SNAPSHOT-jar-with-dependencies.jar com.mycompany.app.App
}

br(){
	$0 build && $0 run
}

$@