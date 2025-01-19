//Here’s a typical structure for a Go project:

go-microservice-starter/
├── cmd/                # Entry points for applications
│   └── main.go         # Main application file
├── pkg/                # Shared libraries and reusable packages
│   └── logger/         # Logging utilities
├── internal/           # Private application code
│   └── service/        # Business logic and services
├── configs/            # Configuration files
├── docs/               # Documentation
├── Dockerfile          # Docker configuration
├── go.mod              # Module dependencies
├── go.sum              # Dependency checksums
└── README.md           # Project description

Level Constant	    Numeric Value	        Description

zapcore.DebugLevel	   -1	                Detailed debug information.
zapcore.InfoLevel	    0	                General operational messages.
zapcore.WarnLevel	    1	                Indications of potential issues.
zapcore.ErrorLevel	    2	                Errors that require attention.
zapcore.DPanicLevel	    3	                Development panic: logs and panics in development.
zapcore.PanicLevel	    4	                Logs and then panics.
zapcore.FatalLevel	    5	                Logs and then exits the application.

If log_level in config.json is set to info, only INFO, WARN, ERROR, and higher logs will be written to the file.
DEBUG logs will be ignored.
