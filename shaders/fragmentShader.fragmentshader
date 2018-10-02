#version 330 core
in vec2 Texture;
out vec4 color;
uniform sampler2D tex;
uniform vec4 colorHue;

void main() {
	color = texture(tex, Texture) * colorHue;
}
