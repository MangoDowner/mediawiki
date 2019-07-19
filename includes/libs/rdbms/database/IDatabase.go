package database

/**
 * @defgroup Database Database
 * This group deals with database interface functions
 * and query specifics/optimisations.
 */

/** @var int Callback triggered immediately due to no active transaction */
const TRIGGER_IDLE = 1
/** @var int Callback triggered by COMMIT */
const TRIGGER_COMMIT = 2
/** @var int Callback triggered by ROLLBACK */
const TRIGGER_ROLLBACK = 3

/** @var string Transaction is requested by regular caller outside of the DB layer */
const TRANSACTION_EXPLICIT = ""
/** @var string Transaction is requested internally via DBO_TRX/startAtomic() */
const TRANSACTION_INTERNAL = "implicit"

/** @var string Atomic section is not cancelable */
const ATOMIC_NOT_CANCELABLE = ""
/** @var string Atomic section is cancelable */
const ATOMIC_CANCELABLE = "cancelable"

/** @var string Commit/rollback is from outside the IDatabase handle and connection manager */
const FLUSHING_ONE = ""
/** @var string Commit/rollback is from the connection manager for the IDatabase handle */
const FLUSHING_ALL_PEERS = "flush"
/** @var string Commit/rollback is from the IDatabase handle internally */
const FLUSHING_INTERNAL = "flush-internal"

/** @var string Do not remember the prior flags */
const REMEMBER_NOTHING = ""
/** @var string Remember the prior flags */
const REMEMBER_PRIOR = "remember"
/** @var string Restore to the prior flag state */
const RESTORE_PRIOR = "prior"
/** @var string Restore to the initial flag state */
const RESTORE_INITIAL = "initial"

/** @var string Estimate total time (RTT, scanning, waiting on locks, applying) */
const ESTIMATE_TOTAL = "total"
/** @var string Estimate time to apply (scanning, applying) */
const ESTIMATE_DB_APPLY = "apply"

/** @var int Combine list with comma delimeters */
const LIST_COMMA = 0
/** @var int Combine list with AND clauses */
const LIST_AND = 1
/** @var int Convert map into a SET clause */
const LIST_SET = 2
/** @var int Treat as field name and do not apply value escaping */
const LIST_NAMES = 3
/** @var int Combine list with OR clauses */
const LIST_OR = 4

/** @var int Enable debug logging */
const DBO_DEBUG = 1
/** @var int Disable query buffering (only one result set can be iterated at a time) */
const DBO_NOBUFFER = 2
/** @var int Ignore query errors (internal use only!) */
const DBO_IGNORE = 4
/** @var int Automatically start a transaction before running a query if none is active */
const DBO_TRX = 8
/** @var int Use DBO_TRX in non-CLI mode */
const DBO_DEFAULT = 16
/** @var int Use DB persistent connections if possible */
const DBO_PERSISTENT = 32
/** @var int DBA session mode mostly for Oracle */
const DBO_SYSDBA = 64
/** @var int Schema file mode mostly for Oracle */
const DBO_DDLMODE = 128
/** @var int Enable SSL/TLS in connection protocol */
const DBO_SSL = 256
/** @var int Enable compression in connection protocol */
const DBO_COMPRESS = 512

/**
 * Basic database interface for live and lazy-loaded relation database handles
 *
 * @note IDatabase and DBConnRef should be updated to reflect any changes
 * @ingroup Database
 */
type IDatabase interface {

}
