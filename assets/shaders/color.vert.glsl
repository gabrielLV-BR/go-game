#version 330 core

layout (location = 0) in vec3 inPosition;

void main() {
	vec3 pos = inPosition;
	pos.x *= 0.5;
	gl_Position = vec4(pos, 1.0);
}
