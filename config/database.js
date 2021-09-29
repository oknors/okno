module.exports = ({ env }) => ({
  defaultConnection: 'default',
  connections: {
    default: {
      connector: 'bookshelf',
      settings: {
        client: 'sqlite',
        filename: env('DATABASE_FILENAME', '/var/db/sqlite/okno.db'),
      },
      options: {
        useNullAsDefault: true,
      },
    },
  },
});
