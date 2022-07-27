export default interface LoginData {
  sessionId: string;
  userId: string;
  isOnboarding: boolean;
  expirationDate: string;
  isCompanyCreating: boolean;
  is2FARequired?: boolean;
  type: string;
}
