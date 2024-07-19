package edu.coursera.distributed;

import java.net.ServerSocket;
import java.net.Socket;

import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.PrintWriter;
import java.io.BufferedReader;
import java.io.File;

/**
 * A basic and very limited implementation of a file server that responds to GET
 * requests from HTTP clients.
 */
public final class FileServer {
    /**
     * Main entrypoint for the basic file server.
     *
     * @param socket Provided socket to accept connections on.
     * @param fs A proxy filesystem to serve files from. See the PCDPFilesystem
     *           class for more detailed documentation of its usage.
     * @throws IOException If an I/O error is detected on the server. This
     *                     should be a fatal error, your file server
     *                     implementation is not expected to ever throw
     *                     IOExceptions during normal operation.
     */
    public void run(final ServerSocket socket, final PCDPFilesystem fs)
            throws IOException {
        /*
         * Enter a spin loop for handling client requests to the provided
         * ServerSocket object.
         */
        while (true) {
            Socket socket_ = socket.accept();

            BufferedReader reader = new BufferedReader(new InputStreamReader(socket_.getInputStream()));

            String line = reader.readLine();

            assert line != null;
            assert line.startsWith("GET");
            
            PCDPPath path = new PCDPPath(line.split(" ")[1]);
            PrintWriter writer = new PrintWriter(socket_.getOutputStream());
            String content = fs.readFile(path);
            if(content != null){
                writer.write("HTTP/1.0 200 OK\r\nServer: FileServer\r\n\r\n");
                writer.write(content + "\r\n");
            } else
                writer.write("HTTP/1.0 404 Not Found\r\nServer: FileServer\r\n\r\n");
            writer.close();
        }
    }
}
