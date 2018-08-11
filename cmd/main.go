package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	vertexShaderSource = `
	#version 410
	in vec3 vp;
	void main() {
		gl_Position = vec4(vp, 1.0);
	}
` + "\x00"

	fragmentShaderSource = `
	#version 410
	out vec4 frag_colour;
	void main() {
		frag_colour = vec4(1, 1, 1, 1.0);
	}
` + "\x00"
)

func main() {
	// args
	// args := os.Args[1:]
	// if len(args) != 1 {
	// 	fmt.Println("Usage:")
	// 	fmt.Println("\tServer:\tnemo --server")
	// 	fmt.Println("\tClient:\tnemo <interface>:<port>")
	// }

	// multiplayer
	// node := node.Node{}
	// if args[0] == "-s" || args[0] == "--server" {
	// 	node.ListenAndServe()
	// } else {
	// 	node.Connect(args[0])
	// }

	// render
	runtime.LockOSThread()

	window, err := InitGLFW()
	defer glfw.Terminate()
	if err != nil {
		panic(err)
	}

	program, err := InitOpenGL()
	if err != nil {
		panic(err)
	}

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program)

		drawSquare()

		glfw.PollEvents()
		window.SwapBuffers()
	}
}

func InitGLFW() (*glfw.Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(500, 500, "hi", nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()
	return window, nil
}

func InitOpenGL() (uint32, error) {
	if err := gl.Init(); err != nil {
		return 0, err
	}

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

var square = []float32{
	-0.1, 0.1, 0,
	-0.1, -0.1, 0,
	0.1, -0.1, 0,

	-0.1, 0.1, 0,
	0.1, 0.1, 0,
	0.1, -0.1, 0,
}

func asVertexArrayObject(points []float32) uint32 {
	var vertexBufferObject uint32
	gl.GenBuffers(1, &vertexBufferObject)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObject)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vertexArrayObject uint32
	gl.GenVertexArrays(1, &vertexArrayObject)
	gl.BindVertexArray(vertexArrayObject)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObject)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vertexArrayObject
}

func drawSquare() {
	vertexArray := asVertexArrayObject(square)
	gl.BindVertexArray(vertexArray)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}
