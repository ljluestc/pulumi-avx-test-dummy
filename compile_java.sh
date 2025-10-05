#!/bin/bash

# Java compilation script for Pulumi Aviatrix Demo
echo "☕ Compiling Java Pulumi Aviatrix Demo..."

# Check if javac is available
if ! command -v javac >/dev/null 2>&1; then
    echo "❌ javac not found. Installing Java development tools..."
    if command -v apt-get >/dev/null 2>&1; then
        sudo apt-get update && sudo apt-get install -y openjdk-11-jdk
    elif command -v yum >/dev/null 2>&1; then
        sudo yum install -y java-11-openjdk-devel
    else
        echo "❌ Cannot install Java automatically. Please install Java JDK manually."
        exit 1
    fi
fi

# Compile Java demo
echo "🔨 Compiling demo_java.java..."
javac demo_java.java

if [ $? -eq 0 ]; then
    echo "✅ Java compilation successful!"
    echo "🚀 Running Java demo..."
    java DemoJava
else
    echo "❌ Java compilation failed!"
    exit 1
fi
