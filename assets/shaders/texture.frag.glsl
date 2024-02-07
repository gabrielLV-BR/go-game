#version 330 core

out vec4 outColor;

in vec3 aNormal;
in vec2 aUV;

uniform vec3 uColor;

void main() {
	outColor = vec4(uColor, 1.0) * vec4(aUV.xy, aNormal.z, 1.0);
}
