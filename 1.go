ackage main

import "fmt"

func main() {
    var countryCapitalMap map[string]string /*�������� */
    countryCapitalMap = make(map[string]string)

    /* map����key - value��,�������Ҷ�Ӧ���׶� */
    countryCapitalMap [ "France" ] = "����"
    countryCapitalMap [ "Italy" ] = "����"
    countryCapitalMap [ "Japan" ] = "����"
    countryCapitalMap [ "India " ] = "�µ���"

    /*ʹ�ü������ͼֵ */
    for country := range countryCapitalMap {
        fmt.Println(country, "�׶���", countryCapitalMap [country])
    }

    /*�鿴Ԫ���ڼ������Ƿ���� */
    capital, ok := countryCapitalMap [ "American" ] /*���ȷ������ʵ��,�����,���򲻴��� */
    /*fmt.Println(capital) */
    /*fmt.Println(ok) */
    if (ok) {
        fmt.Println("American ���׶���", capital)
    } else {
        fmt.Println("American ���׶�������")
    }
}