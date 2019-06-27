/**
 * Include most things that are needed to make MediaWiki work.
 *
 * This file is included by WebStart.php and doMaintenance.php so that both
 * web and maintenance scripts share a final set up phase to include necessary
 * files and create global object variables.
 */
package setup

// Disable MWDebug for command line mode, this prevents MWDebug from eating up
// all the memory from logging SQL queries on maintenance scripts
var WgCommandLineMode bool

