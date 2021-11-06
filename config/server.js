module.exports = ({ env }) => ({
  host: env('HOST', '0.0.0.0'),
  port: env.int('PORT', 11111),
  proxy: true,
  admin: {
    auth: {
      secret: env('ADMIN_JWT_SECRET', '6788c2dd5525743f1b81012a16638bba'),
    },
  },
  cron: {
    enabled: true,
  },
});
