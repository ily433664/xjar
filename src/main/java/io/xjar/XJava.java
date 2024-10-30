package io.xjar;

/**
 * Java 信息
 *
 * @author dhe
 * date 2024/10/22
 */
public class XJava {

    /*public static XJava instance = new XJava(System.getProperty("java.version", "unknown"),
            System.getProperty("java.javaMD5", "unknown"),
            System.getProperty("java.javaSHA1", "unknown"));*/
    public static XJava instance = new XJava("unknown", "unknown", "unknown");

    /**
     * Java 版本
     */
    private String version;

    /**
     * Java MD5 值
     */
    private String javaMD5;

    /**
     * Java SHA1 值
     */
    private String javaSHA1;

    public XJava(String version, String javaMD5, String javaSHA1) {
        this.version = version;
        this.javaMD5 = javaMD5;
        this.javaSHA1 = javaSHA1;
    }

    public static XJava of(String version, String javaMD5, String javaSHA1) {
        return new XJava(version, javaMD5, javaSHA1);
    }

    public String getVersion() {
        return version;
    }

    public void setVersion(String version) {
        this.version = version;
    }

    public static XJava getInstance() {
        return instance;
    }

    public static void setInstance(XJava instance) {
        XJava.instance = instance;
    }

    public String getJavaMD5() {
        return javaMD5;
    }

    public void setJavaMD5(String javaMD5) {
        this.javaMD5 = javaMD5;
    }

    public String getJavaSHA1() {
        return javaSHA1;
    }

    public void setJavaSHA1(String javaSHA1) {
        this.javaSHA1 = javaSHA1;
    }
}
