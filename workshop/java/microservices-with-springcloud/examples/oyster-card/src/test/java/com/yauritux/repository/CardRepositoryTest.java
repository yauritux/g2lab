package com.yauritux.repository;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.Assert.assertNull;

import java.math.BigDecimal;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.boot.test.autoconfigure.orm.jpa.TestEntityManager;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import com.yauritux.Application;
import com.yauritux.model.entity.Card;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@RunWith(SpringRunner.class)
@SpringBootTest(classes = Application.class)
@DataJpaTest(showSql = true)
public class CardRepositoryTest {
	
	private static final String CARD_SN = "1234567890";

	@Autowired
	private TestEntityManager em;
	
	@Autowired
	private CardRepository cardRepository;
	
	@Test
	public void findBySerialNo_existingCard_shouldReturnCard() {
		Card card = new Card();
		card.setSerialNo(CARD_SN);
		card.setBalance(BigDecimal.valueOf(150));
		em.persist(card);
		em.flush();
		
		Card found = cardRepository.findBySerialNo(CARD_SN);
		assertThat(found.getSerialNo()).isEqualTo(card.getSerialNo());
	}
	
	@Test
	public void findBySerialNo_nonExistingCard_nullIsReturned() {
		Card notFound = cardRepository.findBySerialNo(CARD_SN);
		assertNull(notFound);
	}
}
