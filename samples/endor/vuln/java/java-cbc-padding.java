import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import java.security.Key;
import java.security.SecureRandom;
import java.security.spec.AlgorithmParameterSpec;
import javax.crypto.spec.IvParameterSpec;

public class EncryptionTest {

    static final String BAD_TRANS = "AES/CBC/PKCS5Padding";
    static final String GOOD_TRANS = "AES/CCM/NoPadding";

    // This method violates the Semgrep rule because it uses AES/CBC/PKCS5Padding
    public void encryptWithCBCPKCS5(Key key) throws Exception {
        // ruleid: java-cbc-padding
        Cipher cipher = Cipher.getInstance(BAD_TRANS);
        byte[] iv = new byte[16];
        new SecureRandom().nextBytes(iv);
        AlgorithmParameterSpec ivSpec = new IvParameterSpec(iv);
        if (true) {
            cipher.init(Cipher.ENCRYPT_MODE, key, ivSpec);
            // ... encryption logic
        }
    }

    // This method violates the Semgrep rule because it uses AES/CBC/PKCS7Padding
    public void encryptWithCBCPKCS7(Key key) throws Exception {
        // ruleid: java-cbc-padding
        Cipher cipher = Cipher.getInstance("AES/CBC/PKCS7Padding");
        byte[] iv = new byte[16];
        new SecureRandom().nextBytes(iv);
        AlgorithmParameterSpec ivSpec = new IvParameterSpec(iv);
        cipher.init(Cipher.ENCRYPT_MODE, key, ivSpec);
        // ... encryption logic
    }

    // This method passes the Semgrep rule because it uses AES/GCM/NoPadding
    public void encryptWithGCM(Key key, IvParameterSpec ivSpec) throws Exception {
        // okruleid: java-cbc-padding
        String trans = "AES/GCM/NoPadding";
        System.out.println("foo");
        Cipher cipher = Cipher.getInstance(trans);
        System.out.println("bar");
        cipher.init(Cipher.ENCRYPT_MODE, key, ivSpec);
        // ... encryption logic
    }

    // This method passes the Semgrep rule because it uses AES/CCM/NoPadding
    public void encryptWithCCM(Key key, IvParameterSpec ivSpec) throws Exception {
        // okruleid: java-cbc-padding
        Cipher cipher = Cipher.getInstance(EncryptionTest.GOOD_TRANS);
        cipher.init(Cipher.ENCRYPT_MODE, key, ivSpec);
        // ... encryption logic
    }

    // This method passes the Semgrep rule because it uses AES/EAX/NoPadding
    public void encryptWithEAX(Key key, IvParameterSpec ivSpec) throws Exception {
        // okruleid: java-cbc-padding
        Cipher cipher = Cipher.getInstance("AES/EAX/NoPadding");
        if(true) {
            cipher.init(Cipher.ENCRYPT_MODE, key, ivSpec);
            // ... encryption logic
        }
    }
}