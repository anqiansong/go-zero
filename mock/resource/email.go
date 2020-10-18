package resource

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"strings"
	"time"
)

var (
	emailDomain = []string{
		"@yahoo.co.kr", "@mti.gov.na", "@pchome.com.tw", "@otenet.gr", "@yandex.ru", "@cairns.net.au", "@omantel.net.om",
		"@kornet.net", "@msn.com", "@superonline.com", "@netzero.net", "@amet.com.ar", "@daum.net(hanmail.net", "@yahoo.com",
		"@ctimail.com", "@wilnetonline.net", "@tiscali.co.uk", "@veri", "@qatar.net.qa", "@googlemail.com", "@webmail.co.za",
		"@zamnet.zm", "@rediffmail.com", "@vodamail.co.za", "@del", "@scs-net.org", "@mail.sy", "@ntc.net.np", "@itccolp.com.hk",
		"@korea.com", "@libero.it", "@afnet.net", "@twcny.rr.com", "@bigpond.com", "@prodigy.net.mx", "@iway.na", "@mynet.com",
		"@xtra.co.nz", "@net.sy", "@sltnet.lk", "@sourcesexpert.com", "@qq.com", "@africaonline.co.ci", "@ntlworld.com",
		"@rawagegypt.com", "@multi-industrie.de", "@ji-net.com", "@mt.net.mk", "@citechco.net", "@sinos.net", "@iafrica.com",
		"@eim.ae", "@samara.co.zw", "@chinaren.com", "@warwick.net", "@etang.com", "@mail.mn", "@seed.net.tw", "@sohu.com",
		"@hcm.fpt.vn", "@africaonline.co.zw", "@netvision.net.il", "@aviso.ci", "@gionline.com.au", "@club-internet.fr",
		"@biznetvigator.com", "@zahav.net.il", "@mail.hk.com", "@emirates.net.ae", "@live.com", "@hanmail.com/net", "@gmail.com",
		"@topmarkeplg.com.tw", "@ndf.vsnl.net.in", "@nesma.net.sa", "@btinternet.com", "@qualitynet.net", "@aim.com", "@sina.com",
		"@sogou.com", "@netvigator.com", "@indigo.ie", "@excite.com", "@hanafos.com", "@yeah.net", "@comcast.net", "@mindspring.com",
		"@terra.es", "@swe.com.hk", "@candel.co.jp", "@mongol.net", "@tm.net.my", "@poczta.onet.pl", "@mweb.co.zw", "@dnet.net.id",
		"@zol.co.zw", "@spark.net.gr", "@yahoo.co.jp", "@aol.com", "@hcm.vnn.vn", "@kalianet.to", "@fa", "@adsl.loxinfo.com",
		"@yemen.net.ye", "@xxx.meh.es", "@cal", "@sancharnet.in", "@infovia.com.ar", "@wannado.fr", "@mail.ru", "@hknet.com",
		"@vsnl.com", "@cwgsy.net", "@mail.com", "@mos.com.np", "@cs.com", "@hn.vnn.vn", "@cytanet.com.cy", "@caron.se", "@naver.com",
		"@ask.com", "@tom.com", "@citiz.com", "@hongkong.com", "@cyber.net.pk", "@pacific.net.sg", "@westnet.com.au", "@infoclub.com.np",
		"@inbox.com", "@ttnet.net.tr", "@be-local.com", "@eyou.com", "@eunet.at", "@yahoo.com.cn", "@swiszcz.com", "@westnet.com.all",
		"@mondis.com", "@verizon.net", "@xx.org.il", "@eircom.net", "@sbcglobal.net", "@walla.com", "@hotmail.com", "@sotelgui.net.gn",
		"@namibnet.com",
	}
)

func RandomEmail() string {
	emailDomainLength := len(emailDomain)
	if emailDomainLength%2 == 0 {
		emailDomainLength -= 1
	}
	unix := time.Now().Unix()
	nameList := pinyin.LazyConvert(RandPersonNameZh(), nil)
	name := strings.Join(nameList, "")
	if unix%2 == 0 {
		name = strings.ReplaceAll(RandPersonNameEn(), " ", "-")
	}
	name = strings.ToLower(name)
	domain := emailDomain[unix%int64(emailDomainLength)]
	return fmt.Sprintf("%s%s", name, domain)

}
