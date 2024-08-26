module.exports = {
    env: {
      browser: true,
      es2021: true,
      node: true,
    },
    extends: [
      'eslint:recommended',
      'plugin:@typescript-eslint/recommended',
      'plugin:react/recommended',
      'plugin:prettier/recommended', // Optional: For integrating Prettier with ESLint
    ],
    parser: '@typescript-eslint/parser',
    parserOptions: {
      ecmaVersion: 12,
      sourceType: 'module',
      ecmaFeatures: {
        jsx: true,
      },
    },
    plugins: [
      'react',
      '@typescript-eslint',
    ],
    rules: {
      'no-console': 'warn', // Warn on console.log statements
      'no-unused-vars': 'warn', // Warn on unused variables
      'semi': ['error', 'always'], // Enforce semicolons at the end of statements
      'quotes': ['error', 'single'], // Enforce single quotes
      '@typescript-eslint/no-explicit-any': 'warn', // Warn on usage of `any` type
      'react/prop-types': 'off', // Disable prop-types rule (if using TypeScript, for example)
    },
    settings: {
      react: {
        version: 'detect', // Detect the React version
      },
    },
  };