#version 330 core

out vec4 outColor;

in vec2 aUV;
in vec3 aNormal;

uniform vec3 uColor;
uniform sampler2D uTexture0;

void main() {
	outColor = vec4(aNormal, 1.0) * vec4(uColor, 1.0) * texture(uTexture0, aUV);
}
