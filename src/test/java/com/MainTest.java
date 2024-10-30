package com;

import io.xjar.XCryptos;
import io.xjar.XJava;

/**
 * 测试类
 *
 * @author dhe
 * date 2024/10/21
 */
public class MainTest {

    public static void main(String[] args) throws Exception {
        String filePath = "C:\\Users\\dhe\\Desktop\\xjar\\build\\";
        String fileName = "data-app-web-0.0.1-SNAPSHOT.jar";
        String password = "123456";
        String javaVersion = "unknown";
        String javaMD5 = "unknown";
        String javaSHA1 = "unknown";

        XJava.instance = new XJava(javaVersion, javaMD5, javaSHA1);

        XCryptos.encryption()
                .from(filePath + fileName)
                .use(password)
                .include("/com/xxx/data/**/*.class")
                .include("/mapper/**/*Mapper.xml")
                .include("/esmapper/**/*Mapper.xml")
                .exclude("/static/**/*")
                .exclude("/conf/*")
                .to(filePath + "app.jar");
    }

}
