/// URL for accessing the PostrgeSQL database (should contain a schema name in the path)
pub const DB_URL: &str = "DATABASE_URL";
/// Log level configuration for the application. For formatting info, see [env_logger's documentation](https://docs.rs/env_logger/latest/env_logger/#enabling-logging)
pub const LOG_LEVEL: &str = "LOG_LEVEL";

#[cfg(test)]
pub mod test {
    /// URL for accessing the PostgreSQL database during integration tests (should not contain a schema name in the path)
    pub const TEST_DB_URL: &str = "TEST_DB_URL";
}
