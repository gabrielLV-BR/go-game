#version 330 core

layout (location = 0) in vec3 inPosition;
layout (location = 1) in float inId;

out float aId;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProj;

void main() {
	gl_Position = uProj * uView * uModel * vec4(inPosition, 1.0);
	aId = inId;
}
