#version 330 core

#define MAX_TEXTURES 1

out vec4 outColor;

in vec3 aFragPosition;
in vec3 aNormal;
in vec2 aUV;

//

struct PointLight {
	vec3 position;
	vec3 color;
};

struct AmbientLight {
	vec3 color;
};

uniform vec3 uColor;
uniform bool uIsTextured;

uniform AmbientLight uAmbientLight;
uniform PointLight uPointLight;
uniform sampler2D uTexture[MAX_TEXTURES];

void main() {
	vec3 normal = normalize(aNormal);
	vec3 lightDirection = normalize(uPointLight.position - aFragPosition);

	float diffuseImpact = max(dot(normal, lightDirection), 0.0);
	vec3 diffuse = diffuseImpact * uPointLight.color;

	vec3 color = uColor;

	if(uIsTextured) {
		color *= texture(uTexture[0], aUV).rgb;
	}

	vec3 result = (diffuse + uAmbientLight.color) * color;
	outColor = vec4(result, 1.0);
}
