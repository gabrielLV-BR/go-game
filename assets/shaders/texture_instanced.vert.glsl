#version 330 core

layout (location = 0) in vec3 inPosition;
layout (location = 1) in vec2 inUV;
layout (location = 2) in mat4 inInstanceMatrix;

out vec2 aUV;

uniform mat4 uView;
uniform mat4 uProj;

void main() {
	gl_Position = uProj * uView * inInstanceMatrix * vec4(inPosition, 1.0);
	aUV = inUV;
}
