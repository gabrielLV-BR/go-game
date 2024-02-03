#version 330 core

out vec4 outColor;

in vec2 aUV;

uniform vec3 uColor;
uniform sampler2D uTexture0;

void main() {
	outColor = texture(uTexture0, aUV) * vec4(uColor, 1.0);
}
