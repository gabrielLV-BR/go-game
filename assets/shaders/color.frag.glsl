#version 330 core

in float aId;

out vec4 outColor;

uniform vec3 uColor;

float map(float val) {
	return (val + 1.0) / 2.0;
}

void main() {
	float red = map(sin(aId));
	float blue = map(cos(aId));
	outColor = vec4(red, blue, 0.2, 1.0);
}
