module.exports = ({ env }) => ({
  email: {
    provider: 'smtp',
    providerOptions: {
      host: 'smtppro.zoho.eu',
      port: 465,
      secure: true,
      username: 'jorm@okno.rs',
      password: 'HsEE45AkGh8CU3a!',
      rejectUnauthorized: true,
      requireTLS: true,
      connectionTimeout: 1,
    },
    settings: {
      from: 'info@okno.rs',
      replyTo: 'info@okno.rs',
    },
  },
  preview: {
    publicationState: 'preview',
    previewUrl: 'http://localhost:8001/preview/:contentType/:id'
  },
  sentry: {
    dsn: env('SENTRY_DSN'),
    sendMetadata: true,
  }
});
