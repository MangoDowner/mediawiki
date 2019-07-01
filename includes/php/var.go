package php

import "fmt"

/**
 * Outputs or returns a parsable string representation of a variable
 * @link http://php.net/manual/en/function.var-export.php
 * @param mixed $expression <p>
 * The variable you want to export.
 * </p>
 * @param bool $return [optional] <p>
 * If used and set to true, var_export will return
 * the variable representation instead of outputing it.
 * </p>
 * &note.uses-ob;
 * @return mixed the variable representation when the return
 * parameter is used and evaluates to true. Otherwise, this function will
 * return &null;.
 * @since 4.2.0
 * @since 5.0
 */
func VarExport (expression interface{}) string {
	return fmt.Sprintf("%s", expression)
}

