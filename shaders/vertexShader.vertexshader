#version 330 core

in vec3 position;
in vec2 texture;
out vec2 Texture;

uniform mat4 MVP;

void main() {
	gl_Position = MVP * vec4(position, 1);
	Texture = texture;
}