#version 330 core

out vec4 outColor;

in vec3 aNormal;
in vec2 aUV;

void main() {
	outColor = vec4(1.0, aNormal) * vec4(1.0, aUV, 1.0);
}
