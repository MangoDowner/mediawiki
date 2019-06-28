package php

/**
 * Checks if or where headers have been sent
 * @link http://php.net/manual/en/function.headers-sent.php
 * @param string $file [optional] <p>
 * If the optional file and
 * line parameters are set,
 * headers_sent will put the PHP source file name
 * and line number where output started in the file
 * and line variables.
 * </p>
 * @param int $line [optional] <p>
 * The line number where the output started.
 * </p>
 * @return bool headers_sent will return false if no HTTP headers
 * have already been sent or true otherwise.
 * @since 4.0
 * @since 5.0
 */
func HeadersSent() bool {
	//TODO
	return false
}
