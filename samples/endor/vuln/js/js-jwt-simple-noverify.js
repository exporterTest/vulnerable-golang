const jwt = require('jwt-simple');

function insecureDecodeToken(token) {
    try {
        // Decode JWT without verification
        const decoded = jwt.decode(token, null, true); // The third parameter `true` disables signature verification
        console.log("Decoded JWT:", decoded);
        return decoded;
    } catch (error) {
        console.error("Failed to decode token:", error);
        return null;
    }
}
