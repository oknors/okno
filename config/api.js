module.exports = ({ env }) => ({
  responses: {
    privateAttributes: ['_v', 'id', 'created_at'],
  },
  rest: {
    defaultLimit: 99999,
    maxLimit: 99999,
  },
});