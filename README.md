# Background 

I am working on an agent where a user installs my agent, and part of the agent is to collection some information about the machine it is installed on. 
So in the part of the code is checking the languages I care about if they are installed and what version. The agent will have to report back to a REST api
with a json payload so some json encoding has to go on as well. 

# Justification

Now I tried to make a struct for each language I care about that follows a Language interface, getting inspired from this [talk](https://www.goinggo.net/2016/12/developing-a-design-philosophy-in-go.html). Some of the ways I am getting the version is pretty much the same, but for example the node version I had to do some parsing.
Going forward for example getting the java version will be different as well. 

My question is am I doing this right? Following the interface? Or am I just duplicating a lot of code? 
Also how would I do json encoding on this? Do I have to go through every language struct and give them the correct keys?