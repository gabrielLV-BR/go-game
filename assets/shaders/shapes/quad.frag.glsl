#version 330 core

uniform vec3 uColor;

out vec4 aColor;

void main() {
	aColor = vec4(uColor.rg, 1.0, 1.0);
}
