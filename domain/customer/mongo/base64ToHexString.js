function base64ToUuid(base64) {
	// Decode base64 to raw bytes
	const bytes = Buffer.from(base64, 'base64');

	// Convert each byte to a 2-char hex string and join
	const hex = [...bytes].map((b) => b.toString(16).padStart(2, '0')).join('');

	// Format into UUID standard 8-4-4-4-12
	return [
		hex.substring(0, 8),
		hex.substring(8, 12),
		hex.substring(12, 16),
		hex.substring(16, 20),
		hex.substring(20),
	].join('-');
}

console.log(base64ToUuid('EiDM7Ak6TgSMZ4KkV90E7Q=='));
