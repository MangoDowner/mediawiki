package config
/**
 * Interface for configuration instances
 *
 * @since 1.23
 */

type Config interface {
	/**
	 * Get a configuration variable such as "Sitename" or "UploadMaintenance."
	 *
	 * @param string $name Name of configuration option
	 * @return mixed Value configured
	 * @throws ConfigException
	 */
	 Get( name string ) interface{}

	/**
	* Check whether a configuration option is set for the given name
	*
	* @param string $name Name of configuration option
	* @return bool
	* @since 1.24
	*/
	Has( name string ) bool
}
